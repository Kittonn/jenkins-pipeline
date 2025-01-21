pipeline {
    agent {
        label "test-agent"
    }

    environment {
        IMAGE_NAME = "ghcr.io/kittonn/jenkins-pipeline"
        REGISTRY_CREDENTIALS_NAME = "ghcr-credentials"
        REGISTRY_URL = "https://ghcr.io"
    }

    stages {
        stage("Build Docker Image") {
            steps {
                script {
                    docker.build("${IMAGE_NAME}:${BUILD_NUMBER}")
                }
            }
        }

        stage("Push Image to GHCR") {
            steps {
                script {    
                    docker.withRegistry(REGISTRY_URL, REGISTRY_CREDENTIALS_NAME) {
                        docker.image("${IMAGE_NAME}:${BUILD_NUMBER}").push()
                    }
                }
            }
        }

        stage("Deploy") {
            agent { label "pre-prod-agent" }
            environment {
                PORT = credentials('PORT')
            }
            steps {
                // script {
                //     docker.image("${IMAGE_NAME}:${BUILD_NUMBER}").withRun("-p 3000:3000 -e PORT=3000")
                // }
                sh "docker run -dp 3000:3000 -e PORT=3000 ${IMAGE_NAME}:${BUILD_NUMBER}"
            }
        }
    }
}
