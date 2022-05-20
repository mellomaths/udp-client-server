import socket
import json
import logging


def check_int(value):
    if value[0] in ('-', '+'):
        return value[1:].isdigit()
    return value.isdigit()


def get_int_response(value):
    if not isinstance(value, int) and not check_int(value):
        msg = 'Tipo de dado inválido para um número inteiro'
        logging.error(msg)
        return msg
    
    logging.warning('Incrementando o número')
    value = int(value) + 1
    return str(value)


def get_char_response(value):
    if not isinstance(value, str) and len(value) != 1:
        msg = 'Tipo de dado inválido para um caractere'
        logging.error(msg)
        return msg

    logging.warning('Invertendo o case do caractere')
    return value.swapcase()


def get_string_response(value):
    if not isinstance(value, str) and len(value) < 1:
        msg = 'Tipo de dado inválido para uma string'
        logging.error(msg)
        return msg

    logging.warning('Invertendo a string')
    return value[::-1]
        

HOST = '127.0.0.1'
PORT = 9922
SERVER = (HOST, PORT)

s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
s.bind(SERVER)
logging.warning('Server escutando...')

while True:
    data, addr = s.recvfrom(8192)
    data = data.decode('utf-8')
    logging.warning('Recebido: %s', data)

    response = None

    json_data = json.loads(data)
    if not json_data.get('tipo') or not json_data.get('val'):
        logging.error('Formato JSON inválido')
        response = 'Formato JSON inválido'

    tipo = json_data.get('tipo')
    value = json_data.get('val')

    if not response:
        if tipo == 'int':
            logging.warning('Tipo de dado recebido: Inteiro')
            response = get_int_response(value)
        elif tipo == 'char':
            logging.warning('Tipo de dado recebido: Caractere')
            response = get_char_response(value)
        elif tipo == 'string':
            logging.warning('Tipo de dado recebido: String')
            response = get_string_response(value)
        else:
            response = 'Tipo inválido'
    
    logging.warning('Resposta: %s', response)
    s.sendto(response.encode('utf-8'), addr)
