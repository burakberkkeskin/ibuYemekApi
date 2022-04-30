pipeline{
  agent any
  environment {
    dockerhub=credentials('dockerhub')
  }

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

    stage ('docker push') {
      steps {
        sh 'echo $dockerhub_PSW | docker login -u $dockerhub_USR --password-stdin'
        sh 'docker push safderun/ibu-yemek-api:latest'
        mail bcc: '', body: 'IBU Yemek API project docker pushed succesfully!', cc: '', from: 'Jenkins', replyTo: '', subject: 'ibuYemekApi Build', to: 'safderun@proton.me'
      }
    }

    stage('deploy') {
      steps {
        echo 'Deploying...'
      }
    }
  }
}

