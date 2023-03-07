import sys
import base64

import yaml

def read_Functionio() -> dict:
    """Read the FunctionIO from stdin."""
    return yaml.load(sys.stdin.read(), yaml.Loader)


def write_Functionio(Functionio: dict):
    """Write the FunctionIO to stdout and exit."""
    sys.stdout.write(yaml.dump(Functionio))
    sys.exit(0)


def main():
    fnio = read_Functionio()

    #print([x for x in fnio if x['name'] == 'resource7'])
    # [4]['resource']['spec']['forProvider']['manifest']['data']['POSTGRESQL_PASSWORD']
    connstring = 'postgresql://sally:sallyspassword@dbserver.example:5555/userdata?connect_timeout=10&sslmode=require&target_session_attrs=primary'
    # connb64 = base64.b64encode(connstring.encode('ascii'))

    # [x['resource']['spec']['forProvider']['manifest']['data']['POSTGRESQL_PASSWORD'] for x in fnio if x['name'] == 'resource7']

    for elem in fnio['desired']['resources']:
        if elem['name'] == 'resource7':
            # elem['resource']['spec']['forProvider']['manifest']['data']['POSTGRESQL_SINGLE_URL'] = connb64.decode('ascii')
            elem['resource']['spec']['forProvider']['manifest']['stringData']['POSTGRESQL_URL'] = connstring

   # write_Functionio(fnio['desired']['resources'][0])
    write_Functionio(fnio)

main()