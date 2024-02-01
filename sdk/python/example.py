# %%
import logging
from chifra import Chifra

logging.basicConfig(level=logging.DEBUG)
chifra = Chifra()

# %%
# export
result = chifra.export("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")
for k,v in result.items():
    print(f"{k:4}:", end="")
    if not isinstance(v, dict):
        print(f" {v}")
    else:
        print("")
        for k2,v2 in v.items():
            print(f" {k2:4}: {v2}")

# %%
