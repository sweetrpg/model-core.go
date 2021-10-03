# -*- coding: utf-8 -*-
__author__ = "Paul Schifferer <dm@sweetrpg.com>"
"""
"""

from reprlib import recursive_repr
import logging
from datetime import datetime
from sweetrpg_model_core.convert.date import to_datetime


class BaseModel(object):
    """A base model object which includes an 'id' property as well as
    audit fields for created, updated, and deleted times.
    """

    def __init__(self, *args, **kwargs):
        """Create a base model object."""
        logging.debug("args: %s, kwargs: %s", args, kwargs)
        now = datetime.utcnow()  # .isoformat()

        self.id = kwargs.get("_id") or kwargs.get("id")
        self.created_at = to_datetime(kwargs.get("created_at")) or now
        self.updated_at = to_datetime(kwargs.get("updated_at")) or now
        self.deleted_at = to_datetime(kwargs.get("deleted_at"))

    @recursive_repr()
    def __repr__(self):
        attrs = []
        for k, v in self.__dict__.items():
            if k.startswith("__"):
                continue
            # v = getattr(self, k)
            attrs.append("{k}={v}".format(k=k, v=v))
        return f'<Volume({", ".join(attrs)})>'

    def to_dict(self) -> dict:
        """Returns a dictionary representation of the model object.

        Returns:
            dict: A Python dictionary.
        """
        d = {}
        for k in dir(self):
            logging.debug("k: %s, type: %s", k, type(k))
            if k.startswith("__") or k.startswith("to_"):
                continue
            d[k] = getattr(self, k)
        return d


class EmbeddedModel(object):
    """An embedded model object, meant to be contained in another model object. It does not
    contain an identifier or audit fields.
    """

    def __init__(self, *args, **kwargs):
        """Create a base model object."""
        logging.debug("args: %s, kwargs: %s", args, kwargs)
