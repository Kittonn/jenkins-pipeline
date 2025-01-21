pipeline{
    agent {
        label "test-agent"
    }

    environment {
        IMAGE_NAME = 'ghcr.io/kittonn/jenkins-pipeline'
    }

    stages {
        stage("Build Docker Image") {
            steps{
                docker.build('${IMAGE_NAME}:${BUILD_NUMBER}')
            }
        }
    }
}