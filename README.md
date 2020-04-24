# stats-summary

Really simple program to take a pcap file and extract statsd keys and count
them.

Assumes packet capture is solely built of statsd statistics.

Also doesn't check what type of metric it's looking at, if it's the same key,
it gets lumped together.
