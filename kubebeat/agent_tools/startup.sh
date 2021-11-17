gpg --batch --gen-key /usr/share/elastic-agent/key_gen
gpg --batch --no-tty --passphrase abc -b -a /usr/share/elastic-agent/state/data/downloads/osquerybeat-8.0.0-SNAPSHOT-linux-x86_64.tar.gz
# On the tests I did so far this is getting overridden by a fleet-managed elastic agent, and end up in the .bak file
printf 'agent.download:\n   pgpfile: "/root/.gnupg/pubring.gpg"\n' >> /usr/share/elastic-agent/state/elastic-agent.yml
