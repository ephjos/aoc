import argparse
import logging
import pathlib
import requests
import time

logging.basicConfig(format='[%(asctime)s:%(levelname)s] %(message)s', datefmt='%d-%b-%y %H:%M:%S', level=logging.INFO)

def get_input(year, day, failures=0, force_query=False):
    if failures > 1:
        raise Exception(f'Could not get input for year {year} day {day}')

    if failures > 0:
        time.sleep(5)

    inputs_dir = pathlib.Path().home() / ".aoc_inputs"
    inputs_dir.mkdir(parents=True, exist_ok=True)
    dest_file = inputs_dir / f'input_{year}_{day}'

    if not force_query and dest_file.exists():
        logging.info("Query cached, loading from disk!")
        with dest_file.open() as df:
            return df.read()

    url = f'https://adventofcode.com/{year}/day/{day}/input'

    session_file = pathlib.Path().home() / "./.aoc_session"

    with session_file.open() as sf:
        session = sf.read().rstrip()
        cookies = {"session": session}
        resp = requests.get(url=url, cookies=cookies)
        if resp.status_code != 200:
            logging.warning(f'Request failed with status {resp.status_code}, with text "{resp.text.rstrip()}"')
            return get_input(year, day, failures=failures+1, force_query=force_query)
        content = resp.text.rstrip()
        with dest_file.open("w") as df:
            df.write(content)
        return content


