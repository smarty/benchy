# Purpose #
Profiles provide **defaults** which can be overridden, thus a profile fulfills the roles of communicating what the benchmark is for (unit like tests, end-to-end like tests, etc.) and it helps us to quickly define how a benchmark should work.

Controlled defaults are `samples` and `long`. For more information about these, see *Benchy's Functions* in *README.md*.

## Fast ##
Provides defaults more suitable for common and small pieces of code where variability in execution time is minimal.

| **Option** | **Default** |
|------------|-------------|
| Samples    | 3           |
| Long       | off         |

## Medium ##
Provides defaults more suitable for object or package level benchmarks where some variability in runtime might be noticed. These sorts of benchmarks are less common than `fast` benchmarks and, because they take longer to run, should be used a little more sparingly.

| **Option** | **Default** |
|------------|-------------|
| Samples    | 10          |
| Long       | off         |

## Full Metrics ##
Provides defaults more suitable for end-to-end style benchmarks where an entire system can be benchmarked and variance in runtime is likely. These sorts of benchmarks are very expensive to run and should be used sparingly.

| **Option** | **Default** |
|------------|-------------|
| Samples    | 25          |
| Long       | on          |