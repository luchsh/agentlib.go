> Dear Java, of course I still love you!

This project tries to explore the possibility of writing a JVMTI agent in the Go language.

The idea attracts me primarily for two reasons:
- Smaller memory footprint and usually simpler comparing to a Java agent.
- Easy to develop complex logic in one agent (such as running a HTTP server inside this agent).

It is interesting to accommodate more than one GC-based languages in the same process.
These language runtimes have their own background threads, different synchronization mechanims,
and even keep communicating with each other.

# Build
Just run following command from project root directory
```
bash -x build.sh
```
> For now, only Linux and Mac OSX are supported.

Please try following command to test
```
java -agentlib:jnigo=hello_options -version
```

`test_lib.go` is the place where you can add more features to the agent.

# Known issues
Current codebase is just a skeleton which seems to be capable of running some demos.
Need more time to complete the jvmti and jni wrappers.
