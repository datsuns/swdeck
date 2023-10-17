export namespace main {
	
	export class Job {
	    name: string;
	    command: string;
	    type: string;
	
	    static createFrom(source: any = {}) {
	        return new Job(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.command = source["command"];
	        this.type = source["type"];
	    }
	}

}

