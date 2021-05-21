from flask import Flask, request

app = Flask(__name__)
@app.route('/', methods=['GET','POST'])
def job():
    firstname = request.form.get('firstname')
    lastname = request.form.get('lastname')
    gender = request.form.get('gender')
    birthday = request.form.get('birthday')
    pincode = request.form.get('pincode')
    address = request.form.get('address')
    contact = request.form.get('contact_number')
    email_id = request.form.get('email_id')

    # print(firstname, lastname, gender, birthday, pincode, address, contact, email_id)

    return "hi"

if __name__ == "__main__":
    app.run(port=5030)
