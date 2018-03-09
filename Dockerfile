FROM busybox
LABEL maintainer=analytics-platform-tech@digital.justice.gov.uk
ADD ./aws_temporary_credentials /aws_temporary_credentials
CMD ["/aws_temporary_credentials"]

