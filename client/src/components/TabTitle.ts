export function getTitle(tabs: { title: string }[], screen: string) {
	return tabs.filter((val) => val.title.toString() === screen)[0].title.toString();
}
