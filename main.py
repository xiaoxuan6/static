import base64
import datetime
import os
import sys


def main(data: str, filename: str):
    date = datetime.date.today()
    path = r"{}".format(date).replace('-', '/')
    if not os.path.exists(path):
        os.makedirs(path)

    img = base64.b64decode(data)
    filepath = r"{}/{}/{}".format(os.getcwd(), path, filename)
    with open(filepath, mode='wb') as f:
        f.write(img)


if __name__ == '__main__':
    if len(sys.argv) == 3:
        item = list(sys.argv)
        main(item[1], item[2])
    else:
        print("invalid params")
