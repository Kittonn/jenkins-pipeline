pipeline {
  agent any

  stages {
    stage("Checkout Code"){
        steps {
            checkout scm
        }
    }

    stage("Login to ghcr.io") {
      steps { 
          echo "Hello World" 
      }
    }
  }
}
