FROM envoyproxy/envoy:v1.21.2
COPY ./envoy.yml /etc/envoy/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml
RUN chmod go+r /etc/envoy/envoy.yaml
EXPOSE 8080