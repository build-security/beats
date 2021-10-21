package beater

import (
	"fmt"
	"github.com/elastic/beats/v7/kubebeat/config"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/common/kubernetes"
	"github.com/elastic/beats/v7/libbeat/logp"
	"time"
)

// kubebeat configuration.
type kubebeat struct {
	done    chan struct{}
	config  config.Config
	client  beat.Client
	watcher kubernetes.Watcher
}

// New creates an instance of kubebeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	client, err := kubernetes.GetKubernetesClient(c.KubeConfig)
	if err != nil {
		return nil, fmt.Errorf("fail to get k8sclient client: %s", err.Error())
	}

	watchOptions := kubernetes.WatchOptions{
		SyncTimeout: c.Period,
		Namespace:   "kube-system",
	}

	watcher, err := kubernetes.NewWatcher(client, &kubernetes.Event{}, watchOptions, nil)
	if err != nil {
		return nil, fmt.Errorf("fail to init k8sclient watcher: %s", err.Error())
	}

	var Handler EventHandler
	watcher.AddEventHandler(Handler)

	bt := &kubebeat{
		done:    make(chan struct{}),
		config:  c,
		watcher: watcher,
	}
	return bt, nil

}

// Run starts kubebeat.
func (bt *kubebeat) Run(b *beat.Beat) error {
	logp.Info("kubebeat is running! Hit CTRL-C to stop it.")

	var err error
	err = bt.watcher.Start()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}
	//
	//	pods, err := bt.watcher.CoreV1().Pods("kube-system").List(context.TODO(),
	//		metav1.ListOptions{
	//			LabelSelector: "tier=control-plane",
	//		})
	//	timestamp := time.Now()
	//	if err != nil {
	//		logp.Error(fmt.Errorf("error fetching pods data: %v", err))
	//		continue
	//	}
	//
	//	events := make([]beat.Event, len(pods.Items))
	//
	//	for _, item := range pods.Items {
	//
	//		item.SetManagedFields(nil)
	//		item.Status.Reset()
	//
	//		event := beat.Event{
	//			Timestamp: timestamp,
	//			Fields: common.MapStr{
	//				"type": b.Info.Name,
	//				"pod":  item,
	//			},
	//		}
	//		events = append(events, event)
	//	}
	//
	//	bt.client.PublishAll(events)
		logp.Info("Events sent")
	}
}

// Stop stops kubebeat.
func (bt *kubebeat) Stop() {
	bt.watcher.Stop()
	bt.client.Close()
	close(bt.done)
}
