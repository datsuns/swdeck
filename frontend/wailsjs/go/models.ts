export namespace main {
	
	export class Entry {
	    type: string;
	    icon: string;
	
	    static createFrom(source: any = {}) {
	        return new Entry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.icon = source["icon"];
	    }
	}

}

