ldapsearch -x -b "dc=example,dc=org" -H ldap://0.0.0.0 -D cn=admin,dc=example,dc=org -w adminpassword "(&(objectclass=inetOrgPerson)(uid=user01))"
