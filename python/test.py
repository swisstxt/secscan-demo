from django.forms import formsets

try:
    suppressed_exception()
except Exception:
    pass

formsets.formset_factory("vulnerable")
