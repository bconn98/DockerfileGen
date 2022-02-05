FROM archlinux
RUN pacman --noconfirm -Syy\
	vim \
	tmux \
	&& pacman -Scc \
	&& useradd -m base

WORKDIR /home/base