pipeline {
    agent any

    stages {
        stage('run unit tests') {
            steps {
                sh 'make run_unit_tests'
            }
        }

        stage('run api tests') {
            steps {
                sh 'make start_application'
                sh 'make run_api_tests'
            }
        }
    }

    post {
        always {
            sh 'make stop_application'
        }
    }
}