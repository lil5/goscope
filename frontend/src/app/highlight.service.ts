import {Injectable} from '@angular/core';

import 'clipboard';

import 'prismjs';
import 'prismjs/plugins/toolbar/prism-toolbar';
import 'prismjs/plugins/copy-to-clipboard/prism-copy-to-clipboard';
import 'prismjs/components/prism-css';
import 'prismjs/components/prism-javascript';
import 'prismjs/components/prism-java';
import 'prismjs/components/prism-markup';
import 'prismjs/components/prism-typescript';
import 'prismjs/components/prism-sass';
import 'prismjs/components/prism-scss';
import 'prismjs/components/prism-json';

declare var Prism: any;

// @ts-ignore
@Injectable()
export class HighlightService {

  // @ts-ignore
  constructor() {
  }

  highlightAll() {
    Prism.highlightAll();
  }
}
