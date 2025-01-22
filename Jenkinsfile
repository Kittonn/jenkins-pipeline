pipeline {
    agent {
        label "test-agent"
    }

    environment {
        IMAGE_NAME = "ghcr.io/kittonn/jenkins-pipeline"
        REGISTRY_CREDENTIALS_NAME = "ghcr-credentials"
        REGISTRY_URL = "https://ghcr.io"
        APP_NAME = "web-api"
    }

    stages {
        stage("Build and Push Image to GHCR") {
            steps {
                script {
                    docker.withRegistry(REGISTRY_URL, REGISTRY_CREDENTIALS_NAME) {
                        docker.build("${IMAGE_NAME}:${BUILD_NUMBER}").push()
                    }
                }
            }
        }

        stage("Deploy") {
            steps {
                sh returnStatus: true, script: "docker stop ${APP_NAME}"
                sh returnStatus: true, script: "docker rm ${APP_NAME} -f"

                script {
                    docker.withRegistry(REGISTRY_URL, REGISTRY_CREDENTIALS_NAME) {
                        docker.image("${IMAGE_NAME}:${BUILD_NUMBER}").run('-e "PORT=3000" -p 3000:3000 --name ${APP_NAME}')
                    }
                }
            }
        }
    }

    post{
        always{
            sh "docker system prune -af"
        }
    }
}