http_code = '418'

match http_code:
    case '200':
        print('OK')
    case '404':
        print('Not Found')
    case '418':
        print('I am a teapot')
    case _:
        print('Unknown code')
