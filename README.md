# oras-cmd-test

1. Start a local repo: `docker run -d -p 5000:5000 --name registry registry:2`
2. create a sample dir filled with random dirs & files with symlinks:
    ```
    mkdir hello-oras hello-oras/dir-1 hello-oras/dir-1/dir-2 hello-oras/dir-1/dir-3
    echo "hello-0" > hello-oras/hello-0.txt
    echo "hello-1" > hello-oras/dir-1/hello-1.txt
    echo "hello-2" > hello-oras/dir-1/dir-2/hello-2.txt
    echo "hello-3" > hello-oras/dir-1/dir-3/hello-3.txt

    cd hello-oras/dir-1/dir-3; ln -s ../dir-2/hello-2.txt hello-2-symlink.txt; cd ../../..
    cd hello-oras/dir-1/dir-3; ln -s ../dir-2 symlink-dir-2; cd ../../..
    cd hello-oras/dir-1/dir-3; cat hello-2-symlink.txt; cd ../../..
    ```
    ##### Note:
    Creating symlink with absolute path, example: `ln -s $PWD/hello-oras/dir-1/dir-2/hello-2.txt hello-oras/dir-1/dir-3`
    
    results in following error when pulling the aritfact:
    ```
    $ oras pull localhost:5000/hello-oras-test:v1 --output test-out/

    ⠋ [                    ]( 113 kB/s) Downloading hello-oras                                                                  310.00/310  B 100.00%    3ms
    └─ sha256:640ef24cc11c45a33e4df64cdec28d0dce370a924e9102a5ad10d14b8dc23db1
    Error: failed to extract tar to /Users/vs030455/Desktop/golang-projects/test-oras/test-out/hello-oras: "/Users/vs030455/Desktop/golang-projects/test-oras/hello-oras/dir-1/dir-2/hello-2.txt" is outside of "hello-oras"
    ```

### Push using go:

1. Run cmd line wrapper and upload: `go run main.go` (which is hardcoded to push dir hello-oras as localhost:5000/hello-symlink-test:v1)
2. oras pull localhost:5000/hello-symlink-test:v1 --output test-out/

### Push using oras cmdline tool:

Install oras: `brew install oras`
3. Push the dir to local registry using oras: `oras push localhost:5000/hello-symlink-test:v1 hello-oras/`

1. Pull artifacts and check symlink are resolved:
    ```
    oras pull localhost:5000/hello-symlink-test:v1 --output test-out/
    ```
#### Refs:
1. https://medium.com/@Oskarr3/giving-your-charts-a-home-in-docker-registry-d195d08e4eb3
2. https://oras.land/docs/commands/oras_push