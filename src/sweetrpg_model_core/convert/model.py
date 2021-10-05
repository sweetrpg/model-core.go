# -*- coding: utf-8 -*-
__author__ = "Paul Schifferer <dm@sweetrpg.com>"
"""
"""

import logging
import json


def to_model(doc, model_class):
    """Convert a database document to a model instance.

    :param Document doc: The input document to convert. This instance must have a to_json() function that returns
        a dictionary of the document data.
    :param class model_class: The type of model class to convert to. This class must accept kwargs
        in its initializer.
    :returns: An instance of the model class, initialized with the document data.
    """
    logging.debug("doc: %s, model_class: %s", doc, model_class)
    if not hasattr(doc, "to_json"):
        return None
    j = doc.to_json()
    logging.debug("j: %s", j)
    data = json.loads(j)
    logging.debug("data: %s", data)
    model = model_class(**data)
    logging.debug("model: %s", model)

    return model


def to_document(model, document_class):
    """Convert a model object to a database document.

    :param BaseModel model: The model object to convert. This instance must have a to_dict() method that returns
        a dictionary of the model data.
    :param class document_class: The type of the document class to convert to. This class must accept kwargs
        in its initializer.
    :returns: An instance of the document class, initialized with the model data.
    """
    logging.debug("model: %s, document_class: %s", model, document_class)
    if not hasattr(model, "to_dict"):
        return None
    data = model.to_dict()
    logging.debug("data: %s", data)
    doc = document_class(**data)
    logging.debug("doc: %s", doc)

    return doc
