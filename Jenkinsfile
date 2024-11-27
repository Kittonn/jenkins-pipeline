pipeline {
  agent any

  environment {
    IMAGE_NAME = "echo-api"
    GITHUB_CREDENTIALS = credentials('kitton-github')
  }  

  stages {
    stage("Build Docker Image"){
      stage("Set Up Docker Buildx") {
      steps {
        sh "docker buildx create --use || echo 'Buildx builder already exists'"
        sh "docker buildx inspect --bootstrap"
        }
      }

      steps {
        sh """
          docker buildx build \
          --platform linux/amd64,linux/arm64 \
          -t ghcr.io/${GITHUB_CREDENTIALS_USR}/${IMAGE_NAME}:${BUILD_NUMBER} \
          --push .
        """
      }
    }


    stage("Lojgin to GitHub Container Registry") {
      steps { 
        sh "docker login ghcr.io -u ${GITHUB_CREDENTIALS_USR} -p ${GITHUB_CREDENTIALS_PSW}"
      }
    }

    stage("Push Docker Image to ghcr.io") {
      steps {
        sh "docker push ghcr.io/${GITHUB_CREDENTIALS_USR}/${IMAGE_NAME}:${BUILD_NUMBER}"
      } 
    }
  }
}
