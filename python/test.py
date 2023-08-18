import hashlib
from django.forms import formsets
from django.db import connection

try:
    suppressed_exception()
except Exception:
    pass

formsets.formset_factory("vulnerable")

hashlib.md5('invalid').hexdigest()

user = "hackme"
with connection.cursor() as cursor:
    cursor.execute(f"SELECT * FROM users WHERE username = {user}")
    user = cursor.fetchone()
