pipeline {
    agent none

    stages {
        stage('DinD Static Pod with Inline Kubeconfig') {
            agent {
                kubernetes {
                    label 'dind-agent'
                    defaultContainer 'docker'
                    yamlFile 'jenkins/worker-template.yaml'
                }
            }
            environment {
                KUBECONFIG_CONTENT = credentials('jenkins-kubeconfig-base64')
                GIT_USER = 'mesutdmrz'
                GIT_EMAIL = 'mesutdmrz@gmail.com'
            }
            steps {
                container('docker') {
                    script {
                        // Inline kubeconfig oluştur
                        sh '''
                        mkdir -p /tmp/kube
                        echo "$KUBECONFIG_CONTENT" | base64 -d > /tmp/kube/config
                        export KUBECONFIG=/tmp/kube/config
                        '''

                        // Overlay seçimi: main → prod, diğer branch → test
                        def overlay = env.BRANCH_NAME == 'main' ? 'overlays/prod/' : 'overlays/test/'
                        echo "Using overlay: ${overlay}"

                        // Kubernetes apply -k
                        sh "kubectl apply -k ${overlay}"

                        // Git push için token kullanımı
                        withCredentials([string(credentialsId: 'git-token', variable: 'GIT_TOKEN')]) {
                            sh """
                            kustomize build ${overlay} > ${overlay}/generated.yaml
                            git config user.name "${GIT_USER}"
                            git config user.email "${GIT_EMAIL}"
                            git add ${overlay}/generated.yaml
                            git commit -m "Update kustomize generated.yaml from Jenkins" || true
                            git remote set-url origin https://mesutdmrz:${GIT_TOKEN}@github.com/mesutdmrz/jenkins-test.git
                            git push origin ${env.BRANCH_NAME}
                            """
                        }
                    }
                }
            }
        }
    }
}
