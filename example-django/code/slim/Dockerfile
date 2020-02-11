FROM python:3.7-slim
COPY requirements.txt /tmp/

# install dependencies
RUN apt update
RUN apt install -y libpq-dev
RUN pip3 install --no-cache-dir --upgrade -r /tmp/requirements.txt

ADD . /app
# add code to container
WORKDIR /app

CMD ["python", "manage.py", "runserver", "0.0.0.0:8000"]