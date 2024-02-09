<h1 align = "center"> Verificador de E-mails</h1>

Ferramenta para a verificação de domínios de e-mail.

Como é o processo? São avaliados três elementos:

1. MX Record (Mail Exchange) - Um registro DNS que especifica os servidores de e-mail autorizados a receberem e-mails para este domínio.

2. SPF Record (Sender Policy Framework) - Mecanismo de segurança que lista os servidores autorizados a enviar mensagens usando o nome deste domínio.

3. DMARC Record (Domain-based Message Authentication Reporting and Conformance) - Elemento que informa os servidores quando uma mensagem não é autenticada, para que tomem a ação apropriada.

Domínios sem estes elementos possuem graves falhas de segurança.

![Captura de tela de 2024-02-09 13-45-53](https://github.com/KalimaraPeleteiro/EmailVerifier/assets/94702837/3db88ad2-e99d-4775-9f6f-d3e57010845d)

![Captura de tela de 2024-02-09 13-46-11](https://github.com/KalimaraPeleteiro/EmailVerifier/assets/94702837/fdd855b8-59a2-43e4-a609-a7e504a1448a)
