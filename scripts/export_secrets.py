import yaml

with open('./config/local.yaml', 'r') as local_file:
    main_config = yaml.safe_load(local_file)

rabbitmq_config = main_config['rabbitMQ']

with open('./config/docker.yaml', 'r') as docker_file:
    docker_config = yaml.safe_load(docker_file)

with open('./config/db/postgres.yaml', 'r') as config_file:
    postgres_config = yaml.safe_load(config_file)

docker_compose_template = {
    'version': '3',
    'services': {
        'rabbitmq': {
            'image': 'rabbitmq:3.10.7-management',
            'ports': ['15672:15672'],
            'environment': {
                'RABBITMQ_DEFAULT_USER': rabbitmq_config['user'],
                'RABBITMQ_DEFAULT_PASS': rabbitmq_config['password'],
            },
            'volumes': ['$HOME/.docker-conf/rabbitmq/data/:/var/lib/rabbitmq/'],
            'healthcheck': {
                'test': ['CMD', 'rabbitmq-diagnostics', '-q', 'ping'],
                'interval': '30s',
                'timeout': '10s',
                'retries': 3,
                'start_period': '10s',
            }
        },
        'postgres': {
            'image': 'postgres:14.9',
            'ports': ['5432:5432'],
            'environment': {
                'POSTGRES_DB': postgres_config['namedb'],
                'POSTGRES_USER': postgres_config['user'],
                'POSTGRES_PASSWORD': postgres_config['password'],
            },
            'volumes': ['$HOME/.docker-conf/postgres/data:/var/lib/postgresql/data'],
            'healthcheck': {
                'test': [ "CMD-SHELL", f"pg_isready -U {postgres_config['user']} -d {postgres_config['namedb']}" ],
                'interval': '10s',
                'timeout': '5s',
                'retries': 5,
                'start_period': '10s'
            }
        },
        'app': {
            'build': '.',
            'depends_on': ['rabbitmq', 'postgres'],
            'environment': {
                'CONFIG_PATH': '/config/local.yaml',
                'DB_PATH': '/config/db/postgres.yaml',
                'DOCKER_PATH': '/config/docker.yaml',
            },
        }
    }
}


with open('docker-compose.yml', 'w') as yaml_file:
    yaml.dump(docker_compose_template, yaml_file, default_flow_style=False)

