## L2.14

This Go snippet demonstrates [three different implementations of the “or-channel” pattern](https://colobu.com/2018/03/26/channel-patterns/#Or_Channel_%E6%A8%A1%E5%BC%8F), a common concurrency utility that merges multiple “done” channels into a single one. The resulting channel closes as soon as any of the input channels is closed, allowing concurrent goroutines to stop early when one task finishes or fails.

or1 uses a fan-out approach. It launches one goroutine per input channel, each waiting for its channel to close. A sync.Once ensures the output channel is closed only once, even if multiple input channels close almost simultaneously. This version is straightforward and efficient for a moderate number of channels.

or2 uses reflection with reflect.Select to dynamically wait on any number of channels. It constructs a slice of SelectCase values and performs a single reflect.Select call, which unblocks as soon as any case receives a signal. Although concise and elegant, this method is typically slower than direct select statements due to reflection overhead.

or3 uses a recursive divide-and-conquer approach. It recursively splits the list of channels in half, combining them pairwise through nested select statements. This method reduces the number of cases in any single select but introduces more goroutines and recursive calls, making it slightly slower in this example.
