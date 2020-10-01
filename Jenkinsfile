pipeline {

  environment {
    registry = "joelsipayung/cicdjenkinskubego"
    dockerImage = ""
  }

  agent any

  stages {

    stage('Checkout Source') {
      steps {
        git 'https://github.com/joelsipayung/cicdgowothkubernetes.git'
      }
    }

    stage('Build image') {
      steps{
        script {
          dockerImage = docker.build registry + ":$BUILD_NUMBER"
        }
      }
    }

    stage('Push Image') {
      steps{
        script {
          docker.withRegistry( "" ) {
            dockerImage.push()
          }
        }
      }
    }


   stage('Deploy App') {
      steps {
        script {
          kubeconfig(credentialsId: 'mykubeconfig', serverUrl: 'https://usw1.kubesail.com') {
            sh 'kubectl apply -f mytask.yaml'
          }
        }
      }
    }


  }

}
