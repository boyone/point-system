pipeline {
    agent any

    stages {
        stage('run code analysis'){
            steps {
                sh 'make run_code_analysis'
                junit 'golangci-lint-junit-report.xml'
            }
        }

        stage('run unit tests') {
            steps {
                sh 'make run_unit_tests'
                junit 'coverage.xml'
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