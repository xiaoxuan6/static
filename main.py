import argparse
import datetime
import os

import requests


def main(url: str, filename: str):
    date = datetime.date.today()
    path = r"{}".format(date).replace('-', '/')
    if not os.path.exists(path):
        os.makedirs(path)

    response = requests.get(url)
    if response.status_code == 200:
        filepath = r"{}/{}/{}".format(os.getcwd(), path, filename)
        with open(filepath, mode='wb') as f:
            f.write(response.content)


if __name__ == '__main__':
    parse = argparse.ArgumentParser()
    parse.add_argument('--url', type=str, default='https://transfer.sh/QSJlve/2023-03-17212303.png')
    parse.add_argument('--filename', type=str, default='a.png')
    args = parse.parse_args()

    main(args.url, args.filename)
