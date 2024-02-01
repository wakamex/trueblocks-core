import os
import requests
import sys
import logging

class Chifra():
    session = requests.Session()
    opts = {
        "abis": {
            "known": {"hotkey": "-k", "type": "switch"},
            "find": {"hotkey": "-f", "type": "flag"},
            "hint": {"hotkey": "-n", "type": "flag"},
            "encode": {"hotkey": "-e", "type": "flag"},
            "fmt": {"hotkey": "-x", "type": "flag"},
            "verbose:": {"hotkey": "-v", "type": "switch"},
            "help": {"hotkey": "-h", "type": "switch"},
        },
        "blocks": {
            "hashes": {"hotkey": "-e", "type": "switch"},
            "uncles": {"hotkey": "-c", "type": "switch"},
            "traces": {"hotkey": "-t", "type": "switch"},
            "uniq": {"hotkey": "-u", "type": "switch"},
            "flow": {"hotkey": "-f", "type": "flag"},
            "logs": {"hotkey": "-l", "type": "switch"},
            "emitter": {"hotkey": "-m", "type": "flag"},
            "topic": {"hotkey": "-B", "type": "flag"},
            "withdrawals": {"hotkey": "-i", "type": "switch"},
            "articulate": {"hotkey": "-a", "type": "switch"},
            "bigRange": {"hotkey": "-r", "type": "flag"},
            "count": {"hotkey": "-U", "type": "switch"},
            "cacheTxs": {"hotkey": "", "type": "switch"},
            "cacheTraces": {"hotkey": "", "type": "switch"},
            "raw": {"hotkey": "-w", "type": "switch"},
            "cache": {"hotkey": "-o", "type": "switch"},
            "fmt": {"hotkey": "-x", "type": "flag"},
            "verbose:": {"hotkey": "-v", "type": "switch"},
            "help": {"hotkey": "-h", "type": "switch"},
        },
        "chunks": {
            "check": {"hotkey": "-c", "type": "switch"},
            "pin": {"hotkey": "-i", "type": "switch"},
            "publish": {"hotkey": "-p", "type": "switch"},
            "remote": {"hotkey": "-r", "type": "switch"},
            "belongs": {"hotkey": "-b", "type": "flag"},
            "firstBlock": {"hotkey": "-F", "type": "flag"},
            "lastBlock": {"hotkey": "-L", "type": "flag"},
            "maxAddrs": {"hotkey": "-m", "type": "flag"},
            "deep": {"hotkey": "-d", "type": "switch"},
            "rewrite": {"hotkey": "-e", "type": "switch"},
            "count": {"hotkey": "-U", "type": "switch"},
            "sleep": {"hotkey": "-s", "type": "flag"},
            "fmt": {"hotkey": "-x", "type": "flag"},
            "verbose:": {"hotkey": "-v", "type": "switch"},
            "help": {"hotkey": "-h", "type": "switch"},
        },
        "config": {
            "paths": {"hotkey": "-a", "type": "switch"},
            "fmt": {"hotkey": "-x", "type": "flag"},
            "verbose:": {"hotkey": "-v", "type": "switch"},
            "help": {"hotkey": "-h", "type": "switch"},
        },
        "export": {
            "appearances": {"hotkey": "-p", "type": "switch"},
            "receipts": {"hotkey": "-r", "type": "switch"},
            "logs": {"hotkey": "-l", "type": "switch"},
            "traces": {"hotkey": "-t", "type": "switch"},
            "neighbors": {"hotkey": "-n", "type": "switch"},
            "accounting": {"hotkey": "-C", "type": "switch"},
            "statements": {"hotkey": "-A", "type": "switch"},
            "balances": {"hotkey": "-b", "type": "switch"},
            "withdrawals": {"hotkey": "-i", "type": "switch"},
            "articulate": {"hotkey": "-a", "type": "switch"},
            "cacheTraces": {"hotkey": "-R", "type": "switch"},
            "count": {"hotkey": "-U", "type": "switch"},
            "firstRecord": {"hotkey": "-c", "type": "flag"},
            "maxRecords": {"hotkey": "-e", "type": "flag"},
            "relevant": {"hotkey": "-N", "type": "switch"},
            "emitter": {"hotkey": "-m", "type": "flag"},
            "reverted": {"hotkey": "-V", "type": "switch"},
            "topic": {"hotkey": "-B", "type": "flag"},
            "asset": {"hotkey": "-P", "type": "flag"},
            "flow": {"hotkey": "-f", "type": "flag"},
            "factory": {"hotkey": "-y", "type": "switch"},
            "unripe": {"hotkey": "-u", "type": "switch"},
            "reversed": {"hotkey": "-E", "type": "switch"},
            "noZero": {"hotkey": "-z", "type": "switch"},
            "firstBlock": {"hotkey": "-F", "type": "flag"},
            "lastBlock": {"hotkey": "-L", "type": "flag"},
            "ether": {"hotkey": "-H", "type": "switch"},
            "cache": {"hotkey": "-o", "type": "switch"},
            "fmt": {"hotkey": "-x", "type": "flag"},
            "verbose:": {"hotkey": "-v", "type": "switch"},
            "help": {"hotkey": "-h", "type": "switch"},
        },
        "init": {
            "all": {"hotkey": "-a", "type": "switch"},
            "dryRun": {"hotkey": "-d", "type": "switch"},
            "firstBlock": {"hotkey": "-F", "type": "flag"},
            "sleep": {"hotkey": "-s", "type": "flag"},
            "fmt": {"hotkey": "-x", "type": "flag"},
            "verbose:": {"hotkey": "-v", "type": "switch"},
            "help": {"hotkey": "-h", "type": "switch"},
        },
        "list": {
            "count": {"hotkey": "-U", "type": "switch"},
            "noZero": {"hotkey": "-z", "type": "switch"},
            "bounds": {"hotkey": "-b", "type": "switch"},
            "unripe": {"hotkey": "-u", "type": "switch"},
            "silent": {"hotkey": "-s", "type": "switch"},
            "firstRecord": {"hotkey": "-c", "type": "flag"},
            "maxRecords": {"hotkey": "-e", "type": "flag"},
            "reversed": {"hotkey": "-E", "type": "switch"},
            "firstBlock": {"hotkey": "-F", "type": "flag"},
            "lastBlock": {"hotkey": "-L", "type": "flag"},
            "fmt": {"hotkey": "-x", "type": "flag"},
            "verbose:": {"hotkey": "-v", "type": "switch"},
            "help": {"hotkey": "-h", "type": "switch"},
        },
        "logs": {
            "articulate": {"hotkey": "-a", "type": "switch"},
            "raw": {"hotkey": "-w", "type": "switch"},
            "cache": {"hotkey": "-o", "type": "switch"},
            "fmt": {"hotkey": "-x", "type": "flag"},
            "verbose:": {"hotkey": "-v", "type": "switch"},
            "help": {"hotkey": "-h", "type": "switch"},
        },
        "monitors": {
            "delete": {"hotkey": "", "type": "switch"},
            "undelete": {"hotkey": "", "type": "switch"},
            "remove": {"hotkey": "", "type": "switch"},
            "clean": {"hotkey": "-C", "type": "switch"},
            "list": {"hotkey": "-l", "type": "switch"},
            "watch": {"hotkey": "-w", "type": "switch"},
            "watchlist": {"hotkey": "-a", "type": "flag"},
            "commands": {"hotkey": "-c", "type": "flag"},
            "batchSize": {"hotkey": "-b", "type": "flag"},
            "sleep": {"hotkey": "-s", "type": "flag"},
            "fmt": {"hotkey": "-x", "type": "flag"},
            "verbose:": {"hotkey": "-v", "type": "switch"},
            "help": {"hotkey": "-h", "type": "switch"},
        },
        "names": {
            "expand": {"hotkey": "-e", "type": "switch"},
            "matchCase": {"hotkey": "-m", "type": "switch"},
            "all": {"hotkey": "-a", "type": "switch"},
            "custom": {"hotkey": "-c", "type": "switch"},
            "prefund": {"hotkey": "-p", "type": "switch"},
            "addr": {"hotkey": "-s", "type": "switch"},
            "tags": {"hotkey": "-g", "type": "switch"},
            "clean": {"hotkey": "-C", "type": "switch"},
            "regular": {"hotkey": "-r", "type": "switch"},
            "dryRun": {"hotkey": "-d", "type": "switch"},
            "autoname": {"hotkey": "-A", "type": "flag"},
            "create": {"hotkey": "", "type": "switch"},
            "update": {"hotkey": "", "type": "switch"},
            "delete": {"hotkey": "", "type": "switch"},
            "undelete": {"hotkey": "", "type": "switch"},
            "remove": {"hotkey": "", "type": "switch"},
            "fmt": {"hotkey": "-x", "type": "flag"},
            "verbose:": {"hotkey": "-v", "type": "switch"},
            "help": {"hotkey": "-h", "type": "switch"},
        },
        "receipts": {
            "articulate": {"hotkey": "-a", "type": "switch"},
            "raw": {"hotkey": "-w", "type": "switch"},
            "cache": {"hotkey": "-o", "type": "switch"},
            "fmt": {"hotkey": "-x", "type": "flag"},
            "verbose:": {"hotkey": "-v", "type": "switch"},
            "help": {"hotkey": "-h", "type": "switch"},
        },
        "scrape": {
            "blockCnt": {"hotkey": "-n", "type": "flag"},
            "sleep": {"hotkey": "-s", "type": "flag"},
            "touch": {"hotkey": "-l", "type": "flag"},
            "fmt": {"hotkey": "-x", "type": "flag"},
            "verbose:": {"hotkey": "-v", "type": "switch"},
            "help": {"hotkey": "-h", "type": "switch"},
        },
        "slurp": {
            "types": {"hotkey": "-t", "type": "flag"},
            "appearances": {"hotkey": "-p", "type": "switch"},
            "perPage": {"hotkey": "-P", "type": "flag"},
            "sleep": {"hotkey": "-s", "type": "flag"},
            "raw": {"hotkey": "-w", "type": "switch"},
            "cache": {"hotkey": "-o", "type": "switch"},
            "fmt": {"hotkey": "-x", "type": "flag"},
            "verbose:": {"hotkey": "-v", "type": "switch"},
            "help": {"hotkey": "-h", "type": "switch"},
        },
        "state": {
            "parts": {"hotkey": "-p", "type": "flag"},
            "changes": {"hotkey": "-c", "type": "switch"},
            "noZero": {"hotkey": "-z", "type": "switch"},
            "call": {"hotkey": "-l", "type": "flag"},
            "articulate": {"hotkey": "-a", "type": "switch"},
            "proxyFor": {"hotkey": "-r", "type": "flag"},
            "ether": {"hotkey": "-H", "type": "switch"},
            "cache": {"hotkey": "-o", "type": "switch"},
            "fmt": {"hotkey": "-x", "type": "flag"},
            "verbose:": {"hotkey": "-v", "type": "switch"},
            "help": {"hotkey": "-h", "type": "switch"},
        },
        "status": {
            "diagnose": {"hotkey": "-d", "type": "switch"},
            "firstRecord": {"hotkey": "-c", "type": "flag"},
            "maxRecords": {"hotkey": "-e", "type": "flag"},
            "chains": {"hotkey": "-a", "type": "switch"},
            "fmt": {"hotkey": "-x", "type": "flag"},
            "verbose:": {"hotkey": "-v", "type": "switch"},
            "help": {"hotkey": "-h", "type": "switch"},
        },
        "tokens": {
            "parts": {"hotkey": "-p", "type": "flag"},
            "byAcct": {"hotkey": "-b", "type": "switch"},
            "changes": {"hotkey": "-c", "type": "switch"},
            "noZero": {"hotkey": "-z", "type": "switch"},
            "cache": {"hotkey": "-o", "type": "switch"},
            "fmt": {"hotkey": "-x", "type": "flag"},
            "verbose:": {"hotkey": "-v", "type": "switch"},
            "help": {"hotkey": "-h", "type": "switch"},
        },
        "traces": {
            "articulate": {"hotkey": "-a", "type": "switch"},
            "filter": {"hotkey": "-f", "type": "flag"},
            "count": {"hotkey": "-U", "type": "switch"},
            "raw": {"hotkey": "-w", "type": "switch"},
            "cache": {"hotkey": "-o", "type": "switch"},
            "fmt": {"hotkey": "-x", "type": "flag"},
            "verbose:": {"hotkey": "-v", "type": "switch"},
            "help": {"hotkey": "-h", "type": "switch"},
        },
        "transactions": {
            "articulate": {"hotkey": "-a", "type": "switch"},
            "traces": {"hotkey": "-t", "type": "switch"},
            "uniq": {"hotkey": "-u", "type": "switch"},
            "flow": {"hotkey": "-f", "type": "flag"},
            "logs": {"hotkey": "-l", "type": "switch"},
            "emitter": {"hotkey": "-m", "type": "flag"},
            "topic": {"hotkey": "-B", "type": "flag"},
            "accountFor": {"hotkey": "-A", "type": "flag"},
            "cacheTraces": {"hotkey": "", "type": "switch"},
            "ether": {"hotkey": "-H", "type": "switch"},
            "raw": {"hotkey": "-w", "type": "switch"},
            "cache": {"hotkey": "-o", "type": "switch"},
            "fmt": {"hotkey": "-x", "type": "flag"},
            "verbose:": {"hotkey": "-v", "type": "switch"},
            "help": {"hotkey": "-h", "type": "switch"},
        },
        "when": {
            "list": {"hotkey": "-l", "type": "switch"},
            "timestamps": {"hotkey": "-t", "type": "switch"},
            "count": {"hotkey": "-U", "type": "switch"},
            "repair": {"hotkey": "-r", "type": "switch"},
            "check": {"hotkey": "-c", "type": "switch"},
            "update": {"hotkey": "-u", "type": "switch"},
            "deep": {"hotkey": "-d", "type": "switch"},
            "cache": {"hotkey": "-o", "type": "switch"},
            "fmt": {"hotkey": "-x", "type": "flag"},
            "verbose:": {"hotkey": "-v", "type": "switch"},
            "help": {"hotkey": "-h", "type": "switch"},
        }
    }

    def __init__(self):
        self.base_url = "http://localhost:8080/"

    def make_url(self, cmd, posName, defFmt, options, *args):
        fmt = ''
        ret = ''
        skip = False
        for arg in args:
            logging.debug("arg: %s", arg)

            for key in options:
                if arg == options[key]["hotkey"]:
                    arg = f"--{key}"

            if skip is True:
                skip = False
                continue

            val = ''
            option = arg.replace("-","")
            if option == "help":
                os.system(f"chifra {cmd} --help")

            logging.debug("%s:", option)
            if option in options:
                logging.debug("  hotKey: %s", options[option]["hotkey"])
                logging.debug("  type: %s", options[option]["type"])

                match options[option]["type"]:
                    case 'switch':
                        if len(ret) > 0:
                            val += "&"
                        val += option
                    case 'flag':
                        if len(sys.argv) > i+1:
                            if len(ret) > 0:
                                val += "&"
                            val += arg.replace("-", "") + "=" + sys.argv[i+1]
                            if option == "fmt":
                                fmt = sys.argv[i+1]
                            skip = True

            if val == "":
                if len(ret) > 0:
                    val += "&"
                val += arg.replace('-', '') if arg.startswith("-") else f"{posName}={arg}"
            ret += val

        if fmt == '':
            fmt = defFmt
            if len(ret) > 0:
                ret += "&"
            ret += f"fmt={fmt}"

        logging.debug("cmd: %s", cmd)
        logging.debug("fmt: %s", fmt)
        logging.debug("ret: %s", ret)
        return fmt, f"{cmd}?{ret}"

    def query(self, fmt, url):
        return self.session.get(url).json() if fmt == 'json' else self.session.get(url).text

    def is_valid_ommand(self, cmd):
        return cmd in self.opts

    def usage(self, msg):
        print(msg)
        os.system("chifra --help")
        exit(1)

    def abis(self, args):
        fmt, url = self.make_url("abis", "addrs", "json", self.opts["abis"], args)
        return self.query(fmt, self.base_url+url)

    def blocks(self, args):
        fmt, url = self.make_url("blocks", "blocks", "json", self.opts["blocks"], args)
        return self.query(fmt, self.base_url+url)

    def chunks(self, args):
        fmt, url = self.make_url("chunks", "mode", "json", self.opts["chunks"], args)
        return self.query(fmt, self.base_url+url)

    def config(self, args):
        fmt, url = self.make_url("config", "mode", "json", self.opts["config"], args)
        return self.query(fmt, self.base_url+url)

    def export(self, args):
        fmt, url = self.make_url("export", "addrs", "json", self.opts["export"], args)
        return self.query(fmt, self.base_url+url)

    def init(self, args):
        fmt, url = self.make_url("init", "", "json", self.opts["init"], args)
        return self.query(fmt, self.base_url+url)

    def list(self, args):
        fmt, url = self.make_url("list", "addrs", "json", self.opts["list"], args)
        return self.query(fmt, self.base_url+url)

    def logs(self, args):
        fmt, url = self.make_url("logs", "transactions", "json", self.opts["logs"], args)
        return self.query(fmt, self.base_url+url)

    def monitors(self, args):
        fmt, url = self.make_url("monitors", "addrs", "json", self.opts["monitors"], args)
        return self.query(fmt, self.base_url+url)

    def names(self, args):
        fmt, url = self.make_url("names", "terms", "json", self.opts["names"], args)
        return self.query(fmt, self.base_url+url)

    def receipts(self, args):
        fmt, url = self.make_url("receipts", "transactions", "json", self.opts["receipts"], args)
        return self.query(fmt, self.base_url+url)

    def scrape(self, args):
        fmt, url = self.make_url("scrape", "", "json", self.opts["scrape"], args)
        return self.query(fmt, self.base_url+url)

    def slurp(self, args):
        fmt, url = self.make_url("slurp", "addrs", "json", self.opts["slurp"], args)
        return self.query(fmt, self.base_url+url)

    def state(self, args):
        fmt, url = self.make_url("state", "addrs", "json", self.opts["state"], args)
        return self.query(fmt, self.base_url+url)

    def tokens(self, args):
        fmt, url = self.make_url("tokens", "addrs", "json", self.opts["tokens"], args)
        return self.query(fmt, self.base_url+url)

    def traces(self, args):
        fmt, url = self.make_url("traces", "transactions", "json", self.opts["traces"], args)
        return self.query(fmt, self.base_url+url)

    def transactions(self, args):
        fmt, url = self.make_url("transactions", "transactions", "json", self.opts["transactions"], args)
        return self.query(fmt, self.base_url+url)

    def when(self, args):
        fmt, url = self.make_url("when", "blocks", "json", self.opts["when"], args)
        return self.query(fmt, self.base_url+url)

    def status(self, args):
        fmt, url = self.make_url("status", "modes", "json", self.opts["status"], args)
        return self.query(fmt, self.base_url+url)

    def cmdLine(self):
        ret = "chifra"
        for i, arg in enumerate(sys.argv):
            if i > 0:
                ret += (" " + arg)
        return ret
