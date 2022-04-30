pipeline{
  agent any
  stages {
    stage('scm checkout') {

      steps {
        git(credentialsId: 'jenkins-github', url: 'git@github.com:safderun/ibuYemekApi.git', branch: 'master')
      }

    }

    stage('docker build') {
      steps {
        sh 'docker build -t safderun/ibu-yemek-api:latest .'
      }
    }

    stage('deploy') {
      steps {
        echo 'Deploying...'
      }
    }
  }
}

