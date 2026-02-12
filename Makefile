# Makefile na raiz do projeto

.PHONY: dev-connect dev-timer

dev-connect:
	cd apps/connect-server && air

dev-timer:
	cd apps/timer-server && air