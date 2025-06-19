# Benchy #

## Creating a New Benchmark ##
Create a benchmark function such as `BenchmarkSomeStuff(b *testing.B, profile
options.BenchmarkProfile)`. Inside your function, write `benchy.New(b, profile)`
Like below:

```
func BenchmarkSomeStuff(b *testing.B) {
    benchy.New(b, options.Fast)
}
```
*for more information on profiles, **see profiles.md in doc directory***

This will expose a builder using a fluent api *(see 
[fluent api](https://blog.sigplan.org/2021/03/02/fluent-api-practice-and-theory/))*.

## Benchy's Functions ##
**New**: Initializes the builder for Benchy which will eventually be run.

**DontPrintStats**: Turns off stat printing. Stat printing is normally turned on
and will print out a table of benchmark results.

**SetSampleCount**: Sets the number of samples that will be taken. Default is 25
and minimum is 1. `-test.samples n` (where `n` is an integer value) in the CLI
flags takes precedence.

More samples will lead to better statistics, but can take a long time to run.
You can pair this with `-test.benchtime nx` (where `n` is an integer value and
`x` is a unit of time like 'ns' for nanoseconds). By doing this, you can reduce
how long each sample takes to run, default is 1 second.

**ShowMemoryStats**: Turns on the rendering of memory statistics such as memory
growth and allocations per operation.

**RegisterBenchmark**: Adds a new function to be benchmarked. Flags can be set
on a function when registering to cause Benchy to run it differently.

**RegisterSetup**: Adds a setup function to an already registered benchmark.

**RegisterCleanup**: Adds a cleanup function to an already registered benchmark.

**Run**: Runs all the registered benchmarks and returns the results. Results can
be operated on.

## Results ##
Results is a collection of benchmark statistics. Results are instantiated with a
`testing.TB`. When running Benchy from a benchmark function, the returned
results will automatically be loaded with the same `*testing.B` that was set to
Bency. But it can be instantiated with a `*testing.T` to allow for tests to be
operated upon the results in another testing process later on (useful for
regression checks) using `WriteTo` and `ReadFrom`.

### Asserts ###
Calling `AssertThat` on results allows for assertions like `FasterThan` to be
processed on one or more benchmarks.

## Examples ##
Example uses of Benchy can be found in the `example` directory.