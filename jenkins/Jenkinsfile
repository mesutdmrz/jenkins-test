pipeline {
    agent none

    stages {
        stage('DinD Static Pod') {
            agent {
                kubernetes {
                    label 'dind-agent'                     // must match what you want to call this pod
                    defaultContainer 'docker'              // the container to run steps in
                    yamlFile 'jenkins/worker-template.yaml' // <-- points to the YAML in your repo
                }
            }
            steps {
                container('docker') {
                    echo "=== Docker version ==="
                    sh 'docker version'

                    echo "=== List running containers ==="
                    sh 'docker ps'
                    sh 'kubectl --help'
                    sh 'git help -a'

                    echo "=== Build random Docker image ==="
                    sh '''
                    cat <<EOF > Dockerfile
                    FROM alpine:3.18
                    RUN echo "Hello from DinD" > /hello.txt
                    EOF
                    docker build -t random-dind-image .
                    docker images | grep random-dind-image
                    '''
                }
            }
        }
    }
}
