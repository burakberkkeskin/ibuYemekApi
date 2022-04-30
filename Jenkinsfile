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
        mail bcc: '', body: 'IBU Yemek API project built succesfully!!', cc: '', from: '', replyTo: '', subject: 'ibuYemekApi Build', to: 'burakberkkeskin@gmail.com'
      }
    }

    stage('deploy') {
      steps {
        echo 'Deploying...'
      }
    }
  }
}

