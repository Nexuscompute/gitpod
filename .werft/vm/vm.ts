import { exec } from '../util/shell';

import * as Manifests from './manifests'

const KUBECONFIG_PATH = '/mnt/secrets/harvester-kubeconfig/harvester-kubeconfig.yml'

/**
 * Convenience function to kubectl apply a manifest from stdin.
 */
function kubectlApplyManifest(manifest: string, options?: { validate?: boolean }) {
    exec(`
        cat <<EOF | kubectl --kubeconfig ${KUBECONFIG_PATH} apply --validate=${!!options?.validate} -f -
${manifest}
EOF
    `)
}

/**
 * Start a VM
 * Does not wait for the VM to be ready.
 */
export function startVM(options: { name: string }) {
    const namespace = `preview-${options.name}`
    const userDataSecretName = `userdata-${options.name}`

    kubectlApplyManifest(
        Manifests.NamespaceManifest({
            namespace
        })
    )

    kubectlApplyManifest(
        Manifests.VirtualMachineManifest({
            namespace,
            vmName: options.name,
            userDataSecretName
        }),
        { validate: false }
    )

    kubectlApplyManifest(
        Manifests.ServiceManifest({
            vmName: options.name,
            namespace
        })
    )
}

/**
 * Check if a VM with the given name already exists.
 * @returns true if the VM already exists
 */
export function vmExists(options: { name: string }) {
    const namespace = `preview-${options.name}`
    const status = exec(`kubectl --kubeconfig ${KUBECONFIG_PATH} -n ${namespace} get vmi ${options.name}`, { dontCheckRc: true, silent: true })
    return status.code == 0
}

/**
 * Wait until the VM Instance reaches the Running status.
 * If the VM Instance doesn't reach Running before the timeoutMS it will throw an Error.
 */
export function waitForVM(options: { name: string, timeoutMS: number }) {
    const namespace = `preview-${options.name}`
    const startTime = Date.now()
    while (true) {

        const status = exec(`kubectl --kubeconfig ${KUBECONFIG_PATH} -n ${namespace} get vmi ${options.name} -o jsonpath="{.status.phase}"`, { silent: true }).stdout.trim()

        if (status == "Running") {
            return
        }

        const elapsedTimeMs = Date.now() - startTime
        if (elapsedTimeMs > options.timeoutMS) {
            throw new Error("VM didn reach Running status before the timeout")
        }

        console.log(`VM is not yet running. Current status is ${status}. Sleeping 5 seconds`)
        exec('sleep 5', { silent: true })
    }
}

/**
 * Proxy 127.0.0.1:22 to :22 in the VM through the k8s service
 */
export function startSSHProxy(options: { name: string }) {
    const namespace = `preview-${options.name}`
    exec(`sudo kubectl --kubeconfig=${KUBECONFIG_PATH} -n ${namespace} port-forward service/proxy 22:22 > /dev/null &`)
}
