debian:
	wget https://raw.githubusercontent.com/Stavrospanakakis/lfl/main/build/debian/lfl.deb
	sudo dpkg -i ./lfl.deb

arch:
	wget https://raw.githubusercontent.com/Stavrospanakakis/lfl/main/build/arch/PKGBUILD
	makepkg -g >> PKGBUILD && makepkg -sic