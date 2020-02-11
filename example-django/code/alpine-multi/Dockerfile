# Stage 1
FROM python:3.7-alpine AS compile_image
RUN apk update && apk add postgresql-dev gcc libffi-dev musl-dev make


RUN mkdir -p /build/wheels
COPY requirements.txt /tmp/

# Build wheels with system dependencies
RUN pip3 install --upgrade pip && pip3 install wheel
RUN pip3 wheel --wheel-dir=/build/wheels -r /tmp/requirements.txt


# Stage 2
FROM python:3.7-alpine
# Copy wheels and libpq dependencies
COPY --from=compile_image /build/wheels /tmp/wheels
COPY --from=compile_image /usr/lib/lib* /usr/lib/

RUN pip3 install --upgrade pip && pip3 install --no-cache-dir /tmp/wheels/* && rm -rf /tmp/wheels
ADD . /app
# add code to container
WORKDIR /app

CMD ["python", "manage.py", "runserver", "0.0.0.0:8000"]