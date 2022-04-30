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
        sh 'docker build -t safderun/ibu-yemek-api:build .'
      }
    }

    stage ('docker push') {
      steps {

        sh 'docker login -u ${dockerhub.username} -p ${dockerhub.password}'
        script{
          if (env.BRANCH_NAME == 'master') {
            sh 'docker tag safderun/ibu-yemek-api:build safderun/ibu-yemek-api:latest'
            sh 'docker push safderun/ibu-yemek-api:latest'
          } else if (env.BRANCH_NAME == 'dev') {
            sh 'docker tag safderun/ibu-yemek-api:build safderun/ibu-yemek-api:dev'
            sh 'docker push safderun/ibu-yemek-api:dev'
          } else {
            echo 'Unknown branch'
          }
        }
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

