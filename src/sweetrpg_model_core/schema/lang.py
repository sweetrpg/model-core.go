# -*- coding: utf-8 -*-
__author__ = "Paul Schifferer <dm@sweetrpg.com>"
"""Schema for language strings.
"""

from marshmallow.fields import Mapping


class LangString(Mapping):
    """

    """

    mapping_type = dict
