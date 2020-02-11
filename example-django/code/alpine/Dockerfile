# Stage 1
FROM python:3.7-alpine AS compile_image
RUN apk update && apk add postgresql-dev gcc libffi-dev musl-dev make


COPY requirements.txt /tmp/

RUN pip3 install --upgrade pip && pip3 install -r /tmp/requirements.txt
ADD . /app
# add code to container
WORKDIR /app

CMD ["python", "manage.py", "runserver", "0.0.0.0:8000"]