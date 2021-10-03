from setuptools import setup

# Metadata goes in setup.cfg. These are here for GitHub's dependency graph.
setup(
    name="sweetrpg-model-core",
    install_requires=["marshmallow==3.13", "PyMongo[tls,srv]==3.12"],
    extras_require={},
)
