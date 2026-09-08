export namespace pdf {
	
	export class ToolResult {
	    outputPaths: string[];
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new ToolResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.outputPaths = source["outputPaths"];
	        this.message = source["message"];
	    }
	}

}

