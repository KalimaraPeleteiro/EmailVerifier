Ferramenta para a verificação de domínios de e-mail.

Como é o processo?
São avaliados três elementos:

1. MX Record (Mail Exchange)
        Um registro DNS que especifica os servidores de e-mail autorizados a receberem e-mails para este domínio.

2. SPF Record (Sender Policy Framework)
        Mecanismo de segurança que lista os servidores autorizados a enviar mensagens usando o nome deste domínio.

3. DMARC Record (Domain-based Message Authentication Reporting and Conformance)
        Informar os servidores quando uma mensagem não é autenticada, para que tomem a ação apropriada.

Domínios sem estes elementos possuem graves falhas de segurança.