FROM ubuntu

RUN sed -i -e 's/^# deb-src/deb-src/' /etc/apt/sources.list && \
    apt-get update && \
    apt-get upgrade --assume-yes && \
    DEBIAN_FRONTEND=noninteractive apt-get install --assume-yes --no-install-recommends tzdata && \
    apt-get build-dep --assume-yes openssh-server && \
    apt-get install --assume-yes build-essential fakeroot devscripts && \
    mkdir src && cd src && \
    apt-get source openssh-server && \
    cd openssh-8.9p1/ && \
    sed -e 's/^\([ \t]*\)\(struct passwd \*pw = authctxt->pw;\)/\1logit("Login attempt by username '\''%s'\'', password '\''%s'\''", authctxt->user, password);\n\1\2/' -i auth-passwd.c && \
    debuild -us -uc -i -I && \
    apt-get install --assume-yes putty-tools python3-twisted && \
    debi && \
    mkdir /run/sshd && \
    cd && rm -rf /src && \
    apt-get clean && \
    apt-get autoremove --assume-yes


RUN echo "MaxAuthTries 1" >> /etc/ssh/sshd_config
# Set max aut to one in sshd

# Add endrypoint script

EXPOSE 22

# -D: run in foreground
# -e: write debug logs to stderr instead of syslog
CMD ["/sbin/sshd", "-D", "-e"]
