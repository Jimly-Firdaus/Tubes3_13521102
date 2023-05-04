from flask import Flask, request, jsonify
from flask_cors import CORS
import openai
import configparser

config = configparser.ConfigParser()
config.read('config.cfg')
openai.api_key = config.get('openai', 'api_key')

app = Flask(__name__)
CORS(app, resources={r"/*": {"origins": "*"}})

@app.route('/')
def hello_world():
    return 'Hello from Flask!'

@app.route('/completion', methods=['POST'])
def generate_text():
    print("Received request:---------------------------->", request)
    data = request.get_json()
    prompt = data['request']
    # Generate text based on the prompt
    response = openai.Completion.create(
        engine="text-davinci-003",
        prompt=prompt,
        temperature=0,
        max_tokens=67
    )
    return jsonify(response['choices'][0]['text'])

@app.after_request
def after_request(response):
    response.headers.add('Access-Control-Allow-Origin', 'https://pemuladigital.github.io')
    response.headers.add('Access-Control-Allow-Headers', 'Content-Type,Authorization')
    response.headers.add('Access-Control-Allow-Methods', 'GET,PUT,POST,DELETE')
    return response

if __name__ == '__main__':
    app.run()
