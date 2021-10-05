# -*- coding: utf-8 -*-
__author__ = "Paul Schifferer <dm@sweetrpg.com>"
"""
"""

import logging


def convert_document(doc, model_class):
    logging.debug("doc: %s, model_class: %s", doc, model_class)
    data = doc.to_json()
    logging.debug("doc: %s, data: %s", doc, data)
    model = model_class(**data)
    logging.debug("model: %s", model)

    return model
