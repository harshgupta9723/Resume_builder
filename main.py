from flask import Flask, request

app = Flask(__name__)
@app.route('/resume',methods=['POST'])
def resume():
    resume = request.form.get("myForm")
    print(resume)
    return resume

if __name__ == "__main__":
    app.run(port=5030)