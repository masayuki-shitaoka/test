# PythonのFlaskフレームワークを使用した例
from flask import Flask, request
import sqlite3

app = Flask(__name__)

@app.route('/products')
def get_products():
    product_id = request.args.get('id')
    db_connection = sqlite3.connect('database.sqlite')
    cursor = db_connection.cursor()
    query = f"SELECT * FROM products WHERE id = '{product_id}'"
    cursor.execute(query)
    products = cursor.fetchall()
    cursor.close()
    db_connection.close()
    return products

# このコードは、ユーザー入力をそのままSQLクエリに組み込んでいるため、SQLインジェクション攻撃を受けやすい。
# test1
