<!-- Generated with Stardoc: http://skydoc.bazel.build -->

<a name="#cnab"></a>

## cnab

<pre>
cnab(<a href="#cnab-name">name</a>, <a href="#cnab-images">images</a>, <a href="#cnab-invocation_images">invocation_images</a>)
</pre>



### Attributes

<table class="params-table">
  <colgroup>
    <col class="col-param" />
    <col class="col-description" />
  </colgroup>
  <tbody>
    <tr id="cnab-name">
      <td><code>name</code></td>
      <td>
        <a href="https://bazel.build/docs/build-ref.html#name">Name</a>; required
        <p>
          A unique name for this target.
        </p>
      </td>
    </tr>
    <tr id="cnab-images">
      <td><code>images</code></td>
      <td>
        <a href="https://bazel.build/docs/skylark/lib/dict.html">Dictionary: Label -> String</a>; optional
        <p>
          Images that are used by this bundle.
        </p>
      </td>
    </tr>
    <tr id="cnab-invocation_images">
      <td><code>invocation_images</code></td>
      <td>
        <a href="https://bazel.build/docs/build-ref.html#labels">List of labels</a>; required
        <p>
          The array of invocation image definitions for this bundle.
        </p>
      </td>
    </tr>
  </tbody>
</table>


